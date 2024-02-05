import "react"
import "./Main.css"
import {ExpressionBlock} from "../ExpressionBlock/ExpressionBlock"

export const Main = () => {
    return (
        <div className="main_page">
            <header className="header">
                <p>Hello World!</p>
            </header>

            <main className="main">
                <div className="main__content__wrapper">
                    <h1>Main Content</h1>
                </div>
                <ExpressionBlock />
            </main>
            <footer className="footer">
                <p>Footer</p>
            </footer>
        </div>
    )

    
}