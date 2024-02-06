import "react"
import "./ExpressionBlock.css"
import {useState, useEffect} from "react"
import axios from "axios"

export const ExpressionBlock = () => {
    const [expression, setExpression] = useState("")
    const [response, setResponse] = useState({})

    const handleChange = (event) => {
        setExpression(event.target.value)
    }

    const fetchResult = async () => {
        const pattern = /\+/g
        let result = expression.replace(pattern, "PP")
        let url = makeExpressionURL(result)
        console.log(url)
        axios.get(url).then((response) => {
            setResponse(response.data)
            console.log(response)
        })
    }

    return (
        <div className="expression">
            <input
                onChange={handleChange}
                value={expression}
                className="expr_input"
                placeholder="Type here your expression"
            ></input>
            <button onClick={fetchResult} className="expr_button">
                Calculate Expression
            </button>
            <p>Result: {response.Result}</p>
        </div>
    )
}

const generateExpressionID = (length) => {
    return Math.random().toString(36).substring(2, length+2);  
}

const makeExpressionURL = (expression) => {
    let experssionID = generateExpressionID(10)
    let url = `http://localhost:8000/new_expression?value=${expression}&id=${experssionID}`
    return url
}