import "react"
import "./Header.css"


export const Header = () => {
    return (
        <header className="header">
            <div className="header__logo">
                <span className="logo__text">Distributed Calculator</span>
            </div>
            <div className="header__links">
                <ul className="header__menu">
                    <li className="menu__item">
                        <a href="https://github.com/ecol-master/Distributed-Calculator">
                            Github
                        </a>
                    </li>
                    <li className="menu__item">
                        <a href="https://t.me/kuzmindev">Telegram</a>
                    </li>
                </ul>
            </div>
        </header>
    )
}
