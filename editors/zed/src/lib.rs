//! Zed extension for Minecraft Metascript.
//!
//! Highlighting, brackets, indentation and the outline come from the
//! tree-sitter grammar in `editors/tree-sitter-mms`. Everything else
//! (diagnostics, completion, hover, rename, formatting, quick fixes) comes
//! from `mms lsp`, which this extension finds or downloads:
//!
//! 1. `lsp.mms.binary.path` in Zed's settings, if set.
//! 2. `mms` on the worktree's `PATH`.
//! 3. The latest GitHub release of MinecraftMetascript/mms for this
//!    platform, downloaded into the extension's work directory.

use std::fs;

use zed_extension_api::{self as zed, settings::LspSettings, LanguageServerId, Result};

const REPO: &str = "MinecraftMetascript/mms";
const SERVER_NAME: &str = "mms";

struct MmsExtension {
    /// The downloaded binary, once found or fetched in this session.
    cached_binary_path: Option<String>,
}

impl MmsExtension {
    fn server_binary(&mut self, id: &LanguageServerId) -> Result<String> {
        if let Some(path) = &self.cached_binary_path {
            if fs::metadata(path).is_ok_and(|m| m.is_file()) {
                return Ok(path.clone());
            }
        }

        let (os, arch) = zed::current_platform();
        let goos = match os {
            zed::Os::Mac => "darwin",
            zed::Os::Linux => "linux",
            zed::Os::Windows => "windows",
        };
        let goarch =
            match arch {
                zed::Architecture::Aarch64 => "arm64",
                zed::Architecture::X8664 => "amd64",
                zed::Architecture::X86 => return Err(
                    "mms has no 32-bit x86 build; install mms yourself and set lsp.mms.binary.path"
                        .into(),
                ),
            };

        zed::set_language_server_installation_status(
            id,
            &zed::LanguageServerInstallationStatus::CheckingForUpdate,
        );
        let release = match zed::latest_github_release(
            REPO,
            zed::GithubReleaseOptions {
                require_assets: true,
                pre_release: false,
            },
        ) {
            Ok(release) => release,
            Err(err) => {
                // Offline or rate limited: fall back to any earlier download.
                if let Some(path) = find_downloaded_binary(goos) {
                    self.cached_binary_path = Some(path.clone());
                    return Ok(path);
                }
                return Err(format!(
                    "could not find `mms` on PATH or download it from https://github.com/{REPO}/releases ({err})"
                ));
            }
        };

        // Release archives are named mms-<tag>-<goos>-<goarch> and hold a
        // directory of the same name (see .github/workflows/release.yml).
        let stem = format!("mms-{}-{goos}-{goarch}", release.version);
        let (asset_name, file_type) = if goos == "windows" {
            (format!("{stem}.zip"), zed::DownloadedFileType::Zip)
        } else {
            (format!("{stem}.tar.gz"), zed::DownloadedFileType::GzipTar)
        };
        let asset = release
            .assets
            .iter()
            .find(|a| a.name == asset_name)
            .ok_or_else(|| format!("release {} has no asset {asset_name}", release.version))?;

        let version_dir = format!("mms-{}", release.version);
        let binary_path = format!("{version_dir}/{stem}/{}", exe_name(goos));

        if !fs::metadata(&binary_path).is_ok_and(|m| m.is_file()) {
            zed::set_language_server_installation_status(
                id,
                &zed::LanguageServerInstallationStatus::Downloading,
            );
            zed::download_file(&asset.download_url, &version_dir, file_type)
                .map_err(|e| format!("failed to download {asset_name}: {e}"))?;
            zed::make_file_executable(&binary_path)?;

            // Keep only the version just downloaded.
            if let Ok(entries) = fs::read_dir(".") {
                for entry in entries.flatten() {
                    let name = entry.file_name();
                    let name = name.to_string_lossy();
                    if name.starts_with("mms-") && name != version_dir {
                        let _ = fs::remove_dir_all(entry.path());
                    }
                }
            }
        }

        zed::set_language_server_installation_status(
            id,
            &zed::LanguageServerInstallationStatus::None,
        );
        self.cached_binary_path = Some(binary_path.clone());
        Ok(binary_path)
    }
}

fn exe_name(goos: &str) -> &'static str {
    if goos == "windows" {
        "mms.exe"
    } else {
        "mms"
    }
}

/// A binary from an earlier download, if there is one.
fn find_downloaded_binary(goos: &str) -> Option<String> {
    for version in fs::read_dir(".").ok()?.flatten() {
        let version_name = version.file_name().to_string_lossy().into_owned();
        if !version_name.starts_with("mms-") {
            continue;
        }
        for inner in fs::read_dir(version.path()).ok()?.flatten() {
            let inner_name = inner.file_name().to_string_lossy().into_owned();
            let path = format!("{version_name}/{inner_name}/{}", exe_name(goos));
            if fs::metadata(&path).is_ok_and(|m| m.is_file()) {
                return Some(path);
            }
        }
    }
    None
}

impl zed::Extension for MmsExtension {
    fn new() -> Self {
        Self {
            cached_binary_path: None,
        }
    }

    fn language_server_command(
        &mut self,
        id: &LanguageServerId,
        worktree: &zed::Worktree,
    ) -> Result<zed::Command> {
        let binary = LspSettings::for_worktree(SERVER_NAME, worktree)
            .ok()
            .and_then(|s| s.binary);
        let args = binary
            .as_ref()
            .and_then(|b| b.arguments.clone())
            .unwrap_or_else(|| vec!["lsp".into()]);
        let mut env = worktree.shell_env();
        if let Some(extra) = binary.as_ref().and_then(|b| b.env.clone()) {
            env.extend(extra);
        }

        let command = if let Some(path) = binary.and_then(|b| b.path) {
            path
        } else if let Some(path) = worktree.which("mms") {
            path
        } else {
            self.server_binary(id)?
        };

        Ok(zed::Command { command, args, env })
    }

    fn language_server_initialization_options(
        &mut self,
        _id: &LanguageServerId,
        worktree: &zed::Worktree,
    ) -> Result<Option<zed::serde_json::Value>> {
        Ok(LspSettings::for_worktree(SERVER_NAME, worktree)
            .ok()
            .and_then(|s| s.initialization_options))
    }

    fn language_server_workspace_configuration(
        &mut self,
        _id: &LanguageServerId,
        worktree: &zed::Worktree,
    ) -> Result<Option<zed::serde_json::Value>> {
        Ok(LspSettings::for_worktree(SERVER_NAME, worktree)
            .ok()
            .and_then(|s| s.settings))
    }
}

zed::register_extension!(MmsExtension);
