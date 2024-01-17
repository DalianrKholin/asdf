import './App.css'
import { useState, useEffect } from 'react'
import OrganItem from './OrganItem'
import Skeleton, { SkeletonTheme } from 'react-loading-skeleton'
import 'react-loading-skeleton/dist/skeleton.css'
import { useLocation, useNavigate } from 'react-router-dom'
import OrgansInfo from './OrgansInfo'
import AddOrgan from './AddOrgan'

export interface Organ {
    _id: string
    Name: string
    Price: string
    Properties: string
    InStack: number
    [key: string]: string | number
}

export type UserType = 'Admin' | 'User' | null

function App({ cart, setCart }) {
    const URL = 'http://localhost:8080/api/item'
    const [organs, setOrgans] = useState<Organ[]>()
    const [isAdding, setIsAdding] = useState(false)
    const [newOrganData, setNewOrganData] = useState<Omit<Organ, '_id'>>()
    const [token, setToken] = useState('')
    const [isAdmin, setIsAdmin] = useState(false)

    const navigate = useNavigate()

    let location = useLocation()

    useEffect(() => {
        const fetchData = async () => {
            const res = await fetch(URL)
            const data = await res.json()
            setOrgans(data)
        }
        setToken(location.state?.token)
        if (location.state?.userType === 'Admin') {
            setIsAdmin(true)
        }

        fetchData()
    }, [])

    const goToCart = () => {
        navigate('/cart', { state: { cart: cart, token: token } })
    }

    return (
        <SkeletonTheme
            baseColor='#35312f'
            highlightColor='#534d4a'
            borderRadius='10px'
        >
            <h1>Organy na stanie 🫀🫁</h1>
            <div className='card'>
                {organs ? (
                    <OrgansInfo
                        organs={organs}
                        isAdding={isAdding}
                        setIsAdding={setIsAdding}
                        token={token}
                        newOrganData={newOrganData!}
                        isAdmin={isAdmin}
                    />
                ) : (
                    <Skeleton className='skeleton-info' />
                )}

                {isAdding && (
                    <AddOrgan
                        newOrganData={newOrganData!}
                        setNewOrganData={setNewOrganData}
                    />
                )}

                {organs ? (
                    organs.map((organ, idx) => {
                        return (
                            <OrganItem
                                token={token}
                                organ={organ}
                                key={idx}
                                idx={idx}
                                isAdmin={isAdmin}
                                cart={cart}
                                setCart={setCart}
                            />
                        )
                    })
                ) : (
                    <Skeleton count={3} className='skeleton-data' />
                )}
            </div>
            <button onClick={goToCart}>Go to Cart</button>
        </SkeletonTheme>
    )
}

export default App
