import { Organ } from './App'
import './OrganItem.css'

interface Props {
    organ: Organ
    idx: number
}

const OrganItem = ({ organ, idx }: Props) => {
    return (
        <div key={idx} className='organ'>
            {Object.keys(organ).map(key => {
                if (key === 'Id') {
                    return
                }
                return <div>{organ[key]}</div>
            })}
            <div className='buttons'>
                <button>Edit</button>
                <button>Delete</button>
            </div>
        </div>
    )
}

export default OrganItem
