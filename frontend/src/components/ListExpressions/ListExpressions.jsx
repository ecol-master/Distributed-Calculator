import "react"
import "./ListExpressions.css"
import {useState, useEffect} from "react"
import {Expression} from "./Expression/Expression"
import axios from "axios"

export const ListExpressions = () => {
    const [loaded, setLoaded] = useState(false)
    const [exprList, setExprList] = useState([])
    
    const fetchExpressions = () => {
        axios.get("http://localhost:8000/list_of_expressions").then((response) => {
            setExprList(response.data)
        })
    }

    useEffect(() => {
        if (loaded) return
        fetchExpressions()
        setLoaded(true)
    }, [loaded])

    return (
        <>
            <div className="expressions__list">
                <button
                    onClick={() => {
                        fetchExpressions()
                    }}
                >
                    Update Data
                </button>
                <ul>
                    {exprList.map((elem) => {
                        return (
                            <li key={elem.ExpressionID}>
                                <Expression expression={elem} />
                            </li>
                        )
                    })}
                </ul>
            </div>
        </>
    )
}
