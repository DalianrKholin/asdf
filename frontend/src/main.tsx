import React, { useState } from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Cart from './Cart'
import Login from './Login.tsx'

const Main = () => {
    const [cart, setCart] = useState<any>()

    const router = createBrowserRouter([
        {
            path: '/',
            element: <App cart={cart} setCart={setCart} />,
        },
        {
            path: '/cart',
            element: <Cart cart={cart} setCart={setCart} />,
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
