import m from 'mithril'
import styles from '../styles/main.module.scss'

interface domainVerifierResponse {
    isValid: boolean
    hasMX: boolean
    hasSPF: boolean
    hasDMARC: boolean
    spfRecord: string
    dmarcRecrod: string
}

interface IMainView extends m.Component {
    isSent: boolean,
    isSuccess: boolean,
    response: domainVerifierResponse[],
    handleSubmit: (event: any) => void
}

var Main: IMainView = {
    handleSubmit: (event: any): void => {
        event.preventDefault()
        
        const formElement = event.currentTarget.querySelector('form')
        const values = new FormData(formElement)
        const emailAddress = values.get('email')
        fetch('http://localhost:8000/api/verify', {
            method: 'POST',
            body: JSON.stringify([emailAddress])
        })
        .then((res: any) => {
            if (res.status === 200) {
                Main.response = res
                Main.isSent = true
                Main.isSuccess  = true
            } else {
                Main.isSent = true
            }
        })
        .then(() => m.redraw())
    },
    view: () => {
        return m('main', { 
            class: styles.wrapper,
            onsubmit: (event: Event) => Main.handleSubmit(event)
        }, m(
            'form', { class: styles.searchBar }, m(
                'input', { 
                    name: 'email',
                    placeholder: 'joedoe@example.com',
                    class: styles.textField,
                    autocomplete: 'off'
                }
            ),
            m(
                'button', {
                    type: 'submit',
                    class: styles.searchButton
                }, 's'
            )
        ), Main.isSent && m(
            'div', {
                class: Main.isSuccess ? styles.resultSuccess : styles.resultError
            },
            Main.response 
                ? 'it\'s valid e-mail address'
                : 'not a vlid e-mail'
        ))
    },
    isSent: false,
    isSuccess: false,
    response: []
}

export default Main