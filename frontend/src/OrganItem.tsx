import { useState } from 'react'
import { Organ } from './App'
import './OrganItem.css'

interface Props {
    organ: Organ
    idx: number
}

const OrganItem = ({ organ, idx }: Props) => {
    const { Id: _, ...initialOrgansData } = organ
    const [organsData, setOrgansData] = useState(initialOrgansData)
    const [isEditing, setIsEditing] = useState(false)

    const handleDelete = async () => {
        const res = await fetch(
            `http://localhost:8080/admin/itemList?id=${organ.Id}`,
            {
                method: 'DELETE',
            }
        )

        console.log(res)

        if (!res.ok) {
            alert('error')
        } else {
            alert('deleted!')
        }
    }

    const handleUpdate = async () => {
        const res = await fetch(
            `http://localhost:8080/admin/itemList?id=${organ.Id}`,
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(organsData),
            }
        )

        console.log(res)

        if (!res.ok) {
            alert('error')
        } else {
            alert('updated!')
        }
    }

    const handleChange = (key: string, value: string) => {
        setOrgansData(prevData => ({
            ...prevData,
            [key]: value,
        }))
    }

    const handleCancel = () => {
        setOrgansData(initialOrgansData)
        setIsEditing(false)
    }

    return (
        <div key={idx} className='organ'>
            {Object.keys(organ).map(key => {
                if (key === 'Id') {
                    return
                }
                return isEditing ? (
                    <input
                        type='text'
                        value={organsData[key]}
                        onChange={e => handleChange(key, e.target.value)}
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
