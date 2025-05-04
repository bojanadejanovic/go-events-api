<script lang="ts">
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { user } from '$lib/stores/auth';
    import { getEvents, registerForEvent, cancelRegistration } from '$lib/services/api';

    let events: any[] = [];
    let loading = true;
    let error = '';

    onMount(async () => {
        if (!$user) {
            goto('/login');
            return;
        }

        try {
            events = await getEvents();
        } catch (e) {
            error = 'Failed to load events';
        } finally {
            loading = false;
        }
    });

    async function handleRegister(eventId: number) {
        try {
            await registerForEvent(eventId);
            // Refresh events list
            events = await getEvents();
        } catch (e) {
            error = 'Failed to register for event';
        }
    }

    async function handleCancel(eventId: number) {
        try {
            await cancelRegistration(eventId);
            // Refresh events list
            events = await getEvents();
        } catch (e) {
            error = 'Failed to cancel registration';
        }
    }
</script>

<div class="max-w-4xl mx-auto">
    <h1 class="text-3xl font-bold mb-6">All Events</h1>

    {#if error}
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
        </div>
    {/if}

    {#if loading}
        <div class="text-center py-8">Loading events...</div>
    {:else}
        <div class="grid gap-6">
            {#each events as event}
                <div class="bg-white rounded-lg shadow-md p-6">
                    <h2 class="text-xl font-bold mb-2">{event.name}</h2>
                    <p class="text-gray-600 mb-2">{event.description}</p>
                    <p class="text-gray-600 mb-2">
                        <strong>Date:</strong> {new Date(event.date_time).toLocaleString()}
                    </p>
                    <p class="text-gray-600 mb-4">
                        <strong>Location:</strong> {event.location}
                    </p>
                    {#if $user}
                        {#if event.registered}
                            <button
                                on:click={() => handleCancel(event.id)}
                                class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                            >
                                Cancel Registration
                            </button>
                        {:else}
                            <button
                                on:click={() => handleRegister(event.id)}
                                class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
                            >
                                Register
                            </button>
                        {/if}
                    {/if}
                </div>
            {/each}
        </div>
    {/if}
</div> 