import './App.css'
import { useState, useEffect } from 'react'
import OrganItem from './OrganItem'
import Skeleton, { SkeletonTheme } from 'react-loading-skeleton'
import 'react-loading-skeleton/dist/skeleton.css'

export interface Organ {
    Id: string
    Name: string
    Price: string
    Properties: string
    InStack: number
    [key: string]: string | number
}

function App() {
    const URL = 'http://localhost:8080/user/itemList'
    const [organs, setOrgans] = useState<Organ[]>()

    useEffect(() => {
        const fetchData = async () => {
            const res = await fetch(URL)
            const data = await res.json()
            setOrgans(data)
        }

        fetchData()
    }, [])

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
                    </div>
                ) : (
                    <Skeleton className='skeleton-info' />
                )}

                {organs ? (
                    organs.map((organ, idx) => {
                        return <OrganItem organ={organ} key={idx} idx={idx} />
                    })
                ) : (
                    <Skeleton count={3} className='skeleton-data' />
                )}
            </div>
        </SkeletonTheme>
    )
}

export default App
