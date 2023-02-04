import ExpertLogin from './pages/ExpertLogin/ExpertLogin.svelte';
import ExpertPanel from './pages/ExpertPanel/ExpertPanel.svelte';

const routes = {
    '/': ExpertLogin,
    '/expert-panel': ExpertPanel
}

export { routes };