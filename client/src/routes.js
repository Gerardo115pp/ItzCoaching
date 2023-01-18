import HomePage from './pages/Home/Home.svelte';
import ServicesPage from './pages/OurServices/OurServices.svelte';

const routes = {
    '/': HomePage,
    '/servicios': ServicesPage,
}

export { routes };