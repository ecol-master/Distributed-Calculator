import "react"
import "./Main.css"
import {ExpressionBlock} from "../ExpressionBlock/ExpressionBlock"
import {Header} from "../Header/Header"

export const Main = () => {
    return (
        <div className="main_page">
            <div className="main_page_wrapper">
                <Header />
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
        </div>
    )
}
