import './App.css'
import {BrowserRouter} from "react-router-dom";
import history from "./util/history";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {useHotkeys, useLocalStorage} from "@mantine/hooks";
import {ColorSchemeProvider, MantineProvider} from "@mantine/core";
import AppShellComponent from "./components/AppShell.jsx";

export default function App() {
    const [colorScheme, setColorScheme] = useLocalStorage({
        key: 'mantine-color-scheme',
        defaultValue: 'light',
        getInitialValueInEffect: true,
    });

    const toggleColorScheme = (value) =>
        setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

    useHotkeys([['mod+J', () => toggleColorScheme()]]);

    const queryClient = new QueryClient();

    function Content() {
        return (
            <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
                <MantineProvider theme={{colorScheme}}>
                    <AppShellComponent/>
                </MantineProvider>
            </ColorSchemeProvider>
        );
    }

    return (
        <BrowserRouter history={history}>
            <QueryClientProvider client={queryClient}>
                <Content/>
            </QueryClientProvider>
        </BrowserRouter>
    );
}
