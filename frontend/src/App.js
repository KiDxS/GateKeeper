import logo from "./logo.svg";
import "./App.css";

const test = async () => {
  let username = "admin"
  let password = "admin"
  let response = await fetch('http://127.0.0.1:8080/api/v1/user/login', {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    credentials: 'include',
    body: JSON.stringify({
        username,
        password
    })
});

    console.log(response.headers);

    
    response = await fetch('http://127.0.0.1:8080/api/v1/protected', {
      credentials: 'include',
  })
  console.log(response)
};

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" onLoad={test} />
                <p>
                    Edit <code>src/App.js</code> and save to reload.
                </p>
                <a
                    className="App-link"
                    href="https://reactjs.org"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Learn React
                </a>
            </header>
        </div>
    );
}

export default App;
