<script lang="ts">
    import { goto } from '$app/navigation';
    import { signup } from '$lib/services/api';

    let email = '';
    let password = '';
    let error = '';
    let loading = false;

    async function handleSubmit() {
        try {
            loading = true;
            error = '';
            await signup(email, password);
            goto('/login');
        } catch (e) {
            error = 'Failed to create account. Please try again.';
            console.error('Signup error:', e);
        } finally {
            loading = false;
        }
    }
</script>

<div class="max-w-md mx-auto mt-8 p-6 bg-white rounded-lg shadow-md">
    <h2 class="text-2xl font-bold mb-6 text-center pt-4">Sign Up</h2>
    
    {#if error}
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
        </div>
    {/if}

    <form on:submit|preventDefault={handleSubmit} class="space-y-6">
        <div class="text-center">
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input
                type="email"
                id="email"
                bind:value={email}
                required
                class="mt-1 block w-64 mx-auto rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                style="width: 16rem; padding: 0.5rem; border: 1px solid #d1d5db;"
            />
        </div>

        <div class="text-center">
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
                type="password"
                id="password"
                bind:value={password}
                required
                class="mt-1 block w-64 mx-auto rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                style="width: 16rem; padding: 0.5rem; border: 1px solid #d1d5db;"
            />
        </div>

        <div class="text-center" style="margin-top: 1rem; margin-bottom: 1rem;">
            <button
                type="submit"
                disabled={loading}
                class="bg-green-500 hover:bg-green-600 text-white px-6 py-2 rounded transition-colors cursor-pointer"
                style="background-color: #22c55e; color: white; padding: 0.5rem 1.5rem; border-radius: 0.25rem;"
            >
                {loading ? 'Creating account...' : 'Sign Up'}
            </button>
        </div>
    </form>

    <p class="mt-4 text-center">
        Already have an account? <button on:click={() => goto('/login')} class="text-blue-500 hover:text-blue-600 cursor-pointer">Login</button>
    </p>
</div> 