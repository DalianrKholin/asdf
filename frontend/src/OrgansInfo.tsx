import { Organ, UserType } from './App'

interface Props {
    organs: Organ[]
    isAdding: boolean
    setIsAdding: React.Dispatch<React.SetStateAction<boolean>>
    token: string
    newOrganData: Omit<Organ, '_id'>
    isAdmin: boolean
}

const OrgansInfo = ({
    organs,
    isAdding,
    setIsAdding,
    token,
    newOrganData,
    isAdmin,
}: Props) => {
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
        <div className='organ-info'>
            {Object.keys(organs[0]).map(key => {
                if (key === '_id') {
                    return
                }
                return <div>{key}</div>
            })}

            {isAdmin &&
                (!isAdding ? (
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
                ))}
        </div>
    )
}

export default OrgansInfo
