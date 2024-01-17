import { Organ } from './App'

interface Props {
    newOrganData: Omit<Organ, 'Id'>
    setNewOrganData: React.Dispatch<
        React.SetStateAction<Omit<Organ, 'Id'> | undefined>
    >
}

const AddOrgan = ({ newOrganData, setNewOrganData }: Props) => {
    const handleChange = (key: string, value: string | number) => {
        setNewOrganData(prevData => ({
            ...prevData,
            [key]: key === 'Price' || key === 'InStack' ? +value : value,
        }))
    }

    return (
        <div className='organ organ-add'>
            <input
                type='text'
                value={newOrganData?.['Name']}
                onChange={e => handleChange('Name', e.target.value)}
            />
            <input
                type='number'
                value={newOrganData?.['Price']}
                onChange={e => handleChange('Price', e.target.value)}
            />
            <input
                type='text'
                value={newOrganData?.['Properties']}
                onChange={e => handleChange('Properties', e.target.value)}
            />
            <input
                type='number'
                value={newOrganData?.['InStack']}
                onChange={e => handleChange('InStack', e.target.value)}
            />
        </div>
    )
}

export default AddOrgan
