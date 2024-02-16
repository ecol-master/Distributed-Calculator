import "react"
import "./App.css"
import { Header } from "./components/Header/Header";
import { Tabs } from "./components/Tabs/Tabs";

function App() {
  return (
    <div className="main_page">
            <div className="main_page_wrapper">
                <Header />
                <main className="main">
                    <div className="main__content__wrapper">
                        <h1>Взаимодействия с калькулятором</h1>
                        <Tabs />
                    </div>
                </main>
            </div>
        </div>
  )
}

export default App;
