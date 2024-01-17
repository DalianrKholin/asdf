import { useEffect, useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'

const Cart = ({ cart, setCart }) => {
    const navigate = useNavigate()

    let location = useLocation()
    const [token, setToken] = useState()

    useEffect(() => {
        setCart(location.state?.cart)
        setToken(location.state?.token)
    }, [])

    const handleOrder = async () => {
        const res = await fetch('http://localhost:8080/api/order', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                token: token!,
            },
            body: JSON.stringify({
                productInfo: Object.keys(cart).map((item: any) => {
                    return { item: cart[item] }
                }),
            }),
        })

        if (!res.ok) {
            alert('error')
            return
        }

        setCart({})
        alert('wysłano')
        navigate('/', { state: { token: token } })
    }

    return (
        <>
            {cart &&
                Object.keys(cart).map(item => {
                    return (
                        <div>
                            {item}: {cart[item]}
                        </div>
                    )
                })}

            <button onClick={handleOrder}>Order</button>
        </>
    )
}

export default Cart
