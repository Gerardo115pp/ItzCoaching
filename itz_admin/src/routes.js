import AdminLogin from './pages/AdminLogin/AdminLogin.svelte';
import AdminPanel from './pages/AdminPanel/AdminPanel.svelte';

const routes = {
    '/': AdminLogin,
    '/admin-panel': AdminPanel
}

export { routes };