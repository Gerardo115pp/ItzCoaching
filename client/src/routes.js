import HomePage from './pages/Home/Home.svelte';
import ServicesPage from './pages/OurServices/OurServices.svelte';
import ExpertsPage from './pages/OurExperts/OurExperts.svelte';

const routes = {
    '/': HomePage,
    '/servicios': ServicesPage,
    '/expertas': ExpertsPage,
}

export { routes };