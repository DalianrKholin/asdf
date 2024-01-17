import { useState } from 'react'
import './Login.css'
import { useNavigate } from 'react-router-dom'

const Login = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const navigate = useNavigate()

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const res = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        })

        const data = await res.text()
        const json = JSON.parse(data)
        const token = json.token

        console.log(token)

        if (!res.ok) {
            alert('error')
            return
        }

        navigate('/', { state: { token: token } })
    }

    return (
        <form className='login-container' onSubmit={e => handleSubmit(e)}>
            <div className='login-element'>
                <label htmlFor='email'>Email: </label>
                <input
                    type='text'
                    id='email'
                    value={email}
                    onChange={e => setEmail(e.target.value)}
                    required
                />
            </div>
            <div className='login-element'>
                <label htmlFor='login'>Password: </label>
                <input
                    type='password'
                    id='password'
                    value={password}
                    onChange={e => setPassword(e.target.value)}
                    required
                />
            </div>
            <button type='submit'>Login</button>
        </form>
    )
}

export default Login
