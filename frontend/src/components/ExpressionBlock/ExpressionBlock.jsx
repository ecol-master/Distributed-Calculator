import "react"
import "./ExpressionBlock.css"
import {useState} from "react"
import axios from "axios"

export const ExpressionBlock = () => {
    const [expression, setExpression] = useState("")
    const [result, setResult] = useState(0)
    const [response, setResponse] = useState({})
    let {isFetching} = false;

    const handleChange = (event) => {
        setExpression(event.target.value)
        console.log("value is:", event.target.value)
    }

    const fetchResult = () => {
        const pattern = /\+/g
        let result = expression.replace(pattern, "PP")
        let url = "http://localhost:8000/new_expression?value=" + result
        axios.get(url).then((response) => {
            setResponse(response.data)
        }).then(
            isFetching = true
        )
        .then(
            setResult(response.Result)
        ).then(
            isFetching = false
        )
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
            {isFetching ? "Result:" + {result} : "loading"}
            <p>Result: {result}</p>
        </div>
    )
}
