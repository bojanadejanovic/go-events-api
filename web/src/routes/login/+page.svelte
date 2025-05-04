<script lang="ts">
    import { goto } from '$app/navigation';
    import { login } from '$lib/services/api';
    import { setAuth } from '$lib/stores/auth';

    let email = '';
    let password = '';
    let error = '';
    let loading = false;

    async function handleSubmit() {
        try {
            loading = true;
            error = '';
            const response = await login(email, password);
            setAuth(response.user, response.token);
            goto('/');
        } catch (e) {
            error = 'Invalid email or password';
        } finally {
            loading = false;
        }
    }
</script>

<div class="max-w-md mx-auto mt-8 p-6 bg-white rounded-lg shadow-md">
    <h2 class="text-2xl font-bold mb-6 text-center">Login</h2>
    
    {#if error}
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
        </div>
    {/if}

    <form on:submit|preventDefault={handleSubmit} class="space-y-4">
        <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input
                type="email"
                id="email"
                bind:value={email}
                required
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
        </div>

        <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
                type="password"
                id="password"
                bind:value={password}
                required
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
        </div>

        <button
            type="submit"
            disabled={loading}
            class="w-full bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 disabled:bg-blue-300"
        >
            {loading ? 'Logging in...' : 'Login'}
        </button>
    </form>

    <p class="mt-4 text-center">
        Don't have an account? <a href="/signup" class="text-blue-500 hover:text-blue-600">Sign up</a>
    </p>
</div> 