import ExpertsProfilePage from './pages/ExpertProfile/ExpertProfile.svelte'
import ServicesPage from './pages/OurServices/OurServices.svelte';
import ExpertsPage from './pages/OurExperts/OurExperts.svelte';
import HomePage from './pages/Home/Home.svelte';
import SuccessAppointmentPage from './pages/PurchaseFeedback/PurchaseFeedback.svelte'

const routes = {
    '/': HomePage,
    '/servicios': ServicesPage,
    '/expertas': ExpertsPage,
    '/experta/:id': ExpertsProfilePage,
    '/appointment-success': SuccessAppointmentPage
}

export { routes };