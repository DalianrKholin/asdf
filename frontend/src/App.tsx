import './App.css'
import { useState, useEffect } from 'react'
import OrganItem from './OrganItem'
import Skeleton, { SkeletonTheme } from 'react-loading-skeleton'
import 'react-loading-skeleton/dist/skeleton.css'
import { useLocation } from 'react-router-dom'

export interface Organ {
    Id: string
    Name: string
    Price: string
    Properties: string
    InStack: number
    [key: string]: string | number
}

function App() {
    const URL = 'http://localhost:8080/api/item'
    const [organs, setOrgans] = useState<Organ[]>()
    const [isAdding, setIsAdding] = useState(false)
    const [newOrganData, setNewOrganData] = useState<Omit<Organ, 'Id'>>()
    const [token, setToken] = useState('')

    let location = useLocation()

    useEffect(() => {
        const fetchData = async () => {
            const res = await fetch(URL)
            const data = await res.json()
            setOrgans(data)
        }
        setToken(location.state.token)

        fetchData()
    }, [])

    const handleChange = (key: string, value: string | number) => {
        setNewOrganData(prevData => ({
            ...prevData,
            [key]: key === 'Price' || key === 'InStack' ? +value : value,
        }))
    }

    const handleCreate = async () => {
        const res = await fetch(`http://localhost:8080/api/admin/item`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                token: token,
            },
            body: JSON.stringify(newOrganData),
        })

        if (!res.ok) {
            alert('error')
        }

        window.location.reload()
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
                    <div className='organ-info'>
                        {Object.keys(organs[0]).map(key => {
                            if (key === 'Id') {
                                return
                            }
                            return <div>{key}</div>
                        })}
                        {!isAdding ? (
                            <button
                                className='add-btn'
                                onClick={() => setIsAdding(true)}
                            >
                                Create
                            </button>
                        ) : (
                            <div className='buttons'>
                                <button onClick={() => setIsAdding(false)}>
                                    Cancel
                                </button>
                                <button onClick={handleCreate}>Send</button>
                            </div>
                        )}
                    </div>
                ) : (
                    <Skeleton className='skeleton-info' />
                )}

                {isAdding && (
                    <div className='organ organ-add'>
                        <input
                            type='text'
                            value={newOrganData?.['Name']}
                            onChange={e => handleChange('Name', e.target.value)}
                        />
                        <input
                            type='number'
                            value={newOrganData?.['Price']}
                            onChange={e =>
                                handleChange('Price', e.target.value)
                            }
                        />
                        <input
                            type='text'
                            value={newOrganData?.['Properties']}
                            onChange={e =>
                                handleChange('Properties', e.target.value)
                            }
                        />
                        <input
                            type='number'
                            value={newOrganData?.['InStack']}
                            onChange={e =>
                                handleChange('InStack', e.target.value)
                            }
                        />
                    </div>
                )}

                {organs ? (
                    organs.map((organ, idx) => {
                        return (
                            <OrganItem
                                token={token}
                                organ={organ}
                                key={idx}
                                idx={idx}
                            />
                        )
                    })
                ) : (
                    <Skeleton count={3} className='skeleton-data' />
                )}
            </div>
        </SkeletonTheme>
    )
}

export default App
