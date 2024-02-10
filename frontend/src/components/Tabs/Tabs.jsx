import "react"
import {ExpressionBlock} from "../CreateExpression/CreateExpression"
import {ListExpressions} from "../ListExpressions/ListExpressions"
import {ConfigView} from "../ConfigView/ConfigView"
import {useState} from "react"
import "./Tabs.css"

export const Tabs = () => {
    const [activeTab, setActiveTab] = useState(1)

    const getClass = (tabVal) => {
        if (tabVal == activeTab) {
            console.log("here")
            return "active"
        }
        return "hide"
    }

    return (
        <>
            <div className="tabs">
                <ul className="tabs-nav">
                    <li onClick={() => setActiveTab(1)}>
                        <a className={getClass(1)} href="#tab-1">
                            Create Expression
                        </a>
                    </li>
                    <li onClick={() => setActiveTab(2)}>
                        <a className={getClass(2)} href="#tab-2">
                            List of Expressions
                        </a>
                    </li>
                    <li onClick={() => setActiveTab(3)}>
                        <a className={getClass(3)} href="#tab-3">
                            Tab Third
                        </a>
                    </li>
                </ul>
                <div className="tabs-stage">
                    <div className={getClass(1)} id="tab-1">
                        <ExpressionBlock />
                    </div>
                    <div className={getClass(2)} id="tab-2">
                        <ListExpressions />
                    </div>
                    <div className={getClass(3)} id="tab-3">
                        <ConfigView />
                    </div>
                </div>
            </div>
        </>
    )
}
