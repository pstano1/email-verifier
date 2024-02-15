import m from 'mithril'

import Main from './views/main'

m.route(document.body, '/', {
    '/': {
        render: () => {
            return m(Main)
        }
    },
})