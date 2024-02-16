import "react"
import "./Expression.css"
import {useState} from "react"
import axios from "axios"

const ExpressionStatuses = {
    StatusError: "error",
    StatusCalculating: "calculating",
    StatusFinished: "finished",
}

const DoneSvg = () => {
    return (
        <svg
            fill="green"
            xmlns="http://www.w3.org/2000/svg"
            height="32"
            viewBox="0 -960 960 960"
            width="32"
        >
            <path d="M382-240 154-468l57-57 171 171 367-367 57 57-424 424Z" />
        </svg>
    )
}

const ErrorSvg = () => {
    return (
        <svg
            fill="red"
            xmlns="http://www.w3.org/2000/svg"
            height="32"
            viewBox="0 -960 960 960"
            width="32"
        >
            <path d="M480-280q17 0 28.5-11.5T520-320q0-17-11.5-28.5T480-360q-17 0-28.5 11.5T440-320q0 17 11.5 28.5T480-280Zm-40-160h80v-240h-80v240Zm40 360q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z" />
        </svg>
    )
}

const SyncSvg = () => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24">
            <path d="M204-318q-22-38-33-78t-11-82q0-134 93-228t227-94h7l-64-64 56-56 160 160-160 160-56-56 64-64h-7q-100 0-170 70.5T240-478q0 26 6 51t18 49l-60 60ZM481-40 321-200l160-160 56 56-64 64h7q100 0 170-70.5T720-482q0-26-6-51t-18-49l60-60q22 38 33 78t11 82q0 134-93 228t-227 94h-7l64 64-56 56Z" />
        </svg>
    )
}

export const Expression = ({expression}) => {
    const [expr, setNewExpr] = useState(expression)

    const fetchExpression = () => {
        const url = `http://localhost:8000/get_expression?id=${expr.ExpressionID}`
        axios.get(url).then((response) => {
            setNewExpr(response.data)
        })
    }

    const renderResultBlock = () => {
        let text, image
        if (expr.Status == ExpressionStatuses.StatusError) {
            text = "Expression is not Valid"
            image = ErrorSvg()
        }else if (expr.Status == ExpressionStatuses.StatusCalculating) {
            return (
                <>
                    <p>Calculating</p>
                    <button
                        onClick={() => {
                            fetchExpression()
                        }}>
                        {SyncSvg()}
                    </button>
                </>
            )
        }else{
            text = `Result: ${expr.Result}`
            image = DoneSvg()
        }

        return (
            <>
                <p>{text}</p>
                <div>{image}</div>
            </>
        )
    }

    return (
        <>
            <div className="expr__view">
                <div className="expr__view__main">
                    <p>ID: {expression.ExpressionID}</p>
                    <p className="view__value">{expression.Expression}</p>
                </div>
                <div className="expr__view__result">{renderResultBlock()}</div>
            </div>
        </>
    )
}