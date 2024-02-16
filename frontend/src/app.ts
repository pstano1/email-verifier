import m from 'mithril'
import './styles/global.scss'

import Main from './views/main'

m.route(document.body, '/', {
    '/': {
        render: () => {
            return m(Main)
        }
    },
})