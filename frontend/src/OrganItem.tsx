import { useState } from 'react'
import { Organ, UserType } from './App'
import './OrganItem.css'

interface Props {
    organ: Organ
    idx: number
    token: string
    userType: UserType
}

const OrganItem = ({ organ, idx, token, userType }: Props) => {
    const { _id: _, ...initialOrgansData } = organ
    const [organsData, setOrgansData] = useState(initialOrgansData)
    const [isEditing, setIsEditing] = useState(false)

    const handleDelete = async () => {
        const res = await fetch(
            `http://localhost:8080/api/admin/item?id=${organ._id}`,
            {
                method: 'DELETE',
                headers: {
                    token: token,
                },
            }
        )
        if (!res.ok) {
            alert('error')
        }
        window.location.reload()
    }

    const handleUpdate = async () => {
        const res = await fetch(
            `http://localhost:8080/api/admin/item/edit?id=${organ._id}`,
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    token: token,
                },
                body: JSON.stringify(organsData),
            }
        )

        if (!res.ok) {
            alert('error')
        }

        window.location.reload()
    }

    const handleChange = (key: string, value: string | number) => {
        setOrgansData(prevData => ({
            ...prevData,
            [key]: key === 'price' || key === 'inStack' ? +value : value,
        }))
    }

    const handleCancel = () => {
        setOrgansData(initialOrgansData)
        setIsEditing(false)
    }

    return (
        <div key={idx} className='organ'>
            {Object.keys(organ).map(key => {
                if (key === '_id') {
                    return
                }
                return isEditing ? (
                    <input
                        type={
                            key === 'Price' || key === 'InStack'
                                ? 'number'
                                : 'text'
                        }
                        value={organsData[key]}
                        onChange={e => handleChange(key, e.target.value)}
                        className='organ-input'
                    />
                ) : (
                    <div>{organ[key]}</div>
                )
            })}
            <div className='buttons'>
                {!isEditing ? (
                    <>
                        <button onClick={() => setIsEditing(true)}>Edit</button>
                        <button onClick={handleDelete}>Delete</button>
                    </>
                ) : (
                    <>
                        <button onClick={handleCancel}>Cancel</button>
                        <button onClick={handleUpdate}>Update</button>
                    </>
                )}
            </div>
        </div>
    )
}

export default OrganItem
