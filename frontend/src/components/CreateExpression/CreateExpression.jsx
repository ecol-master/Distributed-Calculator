import "react"
import "./CreateExpression.css"
import {useState, useEffect} from "react"
import axios from "axios"

export const ExpressionBlock = () => {
    const [expression, setExpression] = useState("")
    const [response, setResponse] = useState({Result: 0})
    const [loading, setLoading] = useState(false)

    const handleChange = (event) => {
        setExpression(event.target.value)
    }

    const fetchResult = async () => {
        let url = makeExpressionURL(expression)
        console.log(url)
        setLoading(true)
        axios
            .get(url)
            .then((response) => {
                setResponse(response.data)
            })
            .finally(setLoading(false))
    }
    useEffect(() => {
        console.log(loading)
    }, [loading])

    return (
        <div className="expression">
            <div className="experssion__title">
                <h2>Create New Expression</h2>
                <p>Please type your expression with spaces between opearations. My code split expression value by them</p>
            </div>
            <div className="expression__content">
                <div className="experssion__main">
                    <input
                        onChange={handleChange}
                        value={expression}
                        className="expr_input"
                        placeholder="Type here your expression"
                    ></input>
                    <button onClick={fetchResult} className="expr_button">
                        Calculate Expression
                    </button>
                </div>
                <div className="expression__result">
                    <h3>
                    {loading  ? "Loading..."  : `Result: ${response.Result}`}
                    </h3>
                </div>
            </div>
        </div>
    )
}

const generateExpressionID = (length) => {
    return Math.random()
        .toString(36)
        .substring(2, length + 2)
}

const makeExpressionURL = (expression) => {
    let experssionID = generateExpressionID(10)
    let result = expression.replace(/\+/g, "PP")
    result = result.replace(/\(/g, "BO")
    result = result.replace(/\)/g, "BC")
    console.log(result)
    let url = `http://localhost:8000/new_expression?value=${result}&id=${experssionID}`
    // const patternPlus = /\+/g

    return url
}
