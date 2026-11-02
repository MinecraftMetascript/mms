// The upstream work MMS is built on, credited in the editor's About dialog
// and the docs footer. Full license texts ship as third-party-licenses.txt.

export interface Credit {
	name: string;
	url: string;
	by: string;
	byUrl: string;
	/** What MMS uses it for. */
	what: string;
	license?: string;
}

export const deepslate: Credit = {
	name: 'Deepslate',
	url: 'https://github.com/misode/deepslate',
	by: 'Misode',
	byUrl: 'https://github.com/misode',
	what: 'Renders every worldgen preview: noises, density functions, terrain, surface rules and biome maps.',
	license: 'MIT'
};

export const credits: Credit[] = [
	deepslate,
	{
		name: 'mcmeta',
		url: 'https://github.com/misode/mcmeta',
		by: 'Misode',
		byUrl: 'https://github.com/misode',
		what: 'The vanilla worldgen, assets and command tree that previews, book pages and command checks read.'
	},
	{
		name: 'vanilla-mcdoc',
		url: 'https://github.com/SpyglassMC/vanilla-mcdoc',
		by: 'SpyglassMC',
		byUrl: 'https://github.com/SpyglassMC',
		what: 'The data pack schemas that checking, completion and hover docs are generated from.',
		license: 'MIT'
	},
	{
		name: 'CodeMirror',
		url: 'https://codemirror.net',
		by: 'Marijn Haverbeke and contributors',
		byUrl: 'https://github.com/codemirror',
		what: 'The code editor.',
		license: 'MIT'
	}
];
