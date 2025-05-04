<script lang="ts">
    import { user } from '$lib/stores/auth';
    import { goto } from '$app/navigation';
    import '../app.css';
</script>

<nav class="bg-gray-800 text-white p-4">
    <div class="container mx-auto flex justify-between items-center">
        <a href="/" class="text-xl font-bold">📅 Events App</a>
        <div class="space-x-4">
            {#if $user}
                <a href="/events" class="hover:text-gray-300">All Events</a>
                <a href="/my-events" class="hover:text-gray-300">My Events</a>
                <a href="/logout" class="hover:text-gray-300">Logout</a>
            {:else}
                <button 
                    on:click={() => goto('/login')}
                    class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded transition-colors"
                    style="background-color: #3b82f6; color: white; padding: 0.5rem 1rem; border-radius: 0.25rem;"
                >
                    Login
                </button>
                <button 
                    on:click={() => goto('/signup')}
                    class="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded transition-colors"
                    style="background-color: #22c55e; color: white; padding: 0.5rem 1rem; border-radius: 0.25rem;"
                >
                    Sign Up
                </button>
            {/if}
        </div>
    </div>
</nav>

<main class="container mx-auto p-4">
    <slot />
</main>

<style>
    :global(body) {
        margin: 0;
        font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    }
</style> 