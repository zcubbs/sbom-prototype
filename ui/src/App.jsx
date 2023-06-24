import {useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {BrowserRouter} from "react-router-dom";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

function App() {
    const queryClient = new QueryClient()

    const [colorScheme, setColorScheme] = useLocalStorage({
        key: 'mantine-color-scheme',
        defaultValue: 'light',
        getInitialValueInEffect: true,
    });

    return (
        <>
            <BrowserRouter history={history}>
                <QueryClientProvider client={queryClient}>
                    <Content/>
                </QueryClientProvider>
            </BrowserRouter>
        </>
    )
}

function Content() {
    const { isLoading, error, data } = useQuery({
        queryKey: ['repoData'],
        queryFn: () =>
            fetch('https://api.github.com/repos/tannerlinsley/react-query').then(
                (res) => res.json(),
            ),
    })

    if (isLoading) return 'Loading...'
    if (error) return 'An error has occurred: ' + error.message

    return (
        <div className="App"></div>
    )
}

export default App
