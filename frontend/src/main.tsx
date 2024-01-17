import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Cart from './Cart'
import Login from './Login.tsx'

const Main = () => {
    const router = createBrowserRouter([
        {
            path: '/',
            element: <App />,
        },
        {
            path: '/cart',
            element: <Cart />,
        },
        {
            path: '/login',
            element: <Login />,
        },
    ])

    return (
        <React.StrictMode>
            <RouterProvider router={router} />
        </React.StrictMode>
    )
}

export default Main

ReactDOM.createRoot(document.getElementById('root')!).render(<Main />)
