import "react"
import "./ExpressionBlock.css"
import {useState} from "react"

export const ExpressionBlock = () => {
    const [message, setMessage] = useState("")
    const [result, setResult] = useState(0);
    const handleChange = (event) => {
        setMessage(event.target.value)

        console.log("value is:", event.target.value)
    }

    const fetchResult = () => {
        let url = ""
    }

    return (
        <div className="expression">
            <input
                onChange={handleChange}
                value={message}
                className="expr_input"
                placeholder="Type here your expression"
            ></input>
            <button onClick={fetchResult} className="expr_button">
                Calculate Expression
            </button>
            <p>
                Result: {result}
            </p>
        </div>
    )
}
