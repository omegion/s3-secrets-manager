const {description} = require('../../package.json')

module.exports = {
    // site config
    lang: 'en-US',
    title: 'S3 Secrets Manager',
    description: description,

    head: [
        [
            "link",
            {
                rel: "apple-touch-icon",
                sizes: "192x192",
                href: "/favicons/apple-icon-192x192.png",
            },
        ],
        [
            "link",
            {
                rel: "icon",
                type: "image/png",
                sizes: "32x32",
                href: "/favicons/favicon-32x32.png",
            },
        ],
        [
            "link",
            {
                rel: "icon",
                type: "image/png",
                sizes: "16x16",
                href: "/favicons/favicon-16x16.png",
            },
        ],
        ["link", {rel: "shortcut icon", href: "/favicons/favicon.ico"}],
        ["meta", {name: "theme-color", content: "#0842ba"}],
        ["meta", {name: "apple-mobile-web-app-capable", content: "yes"}],
        [
            "meta",
            {name: "apple-mobile-web-app-status-bar-style", content: "black"},
        ],
    ],

    // theme and its config
    theme: '@vuepress/theme-default',
    themeConfig: {
        repo: 'https://github.com/omegion/s3-secrets-manager',
        editLinks: false,
        docsDir: '',
        editLinkText: '',
        lastUpdated: false,
        logo: "/img/logo.svg",
        author: "omegion",
        navbar: [
            {
                text: 'Guide',
                link: '/guide/',
            },
        ],
        sidebar: {
            '/guide/': [
                {
                    title: 'Guide',
                    collapsable: false,
                    children: [
                        'README.md',
                        'get-started.md',
                        'quick-start.md',
                    ]
                }
            ],
        },
    },
    plugins: [
        [
            '@vuepress/plugin-palette',
            { preset: 'sass' },
        ],
    ],
}
