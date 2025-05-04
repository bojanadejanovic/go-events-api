<script lang="ts">
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { clearAuth } from '$lib/stores/auth';

    onMount(() => {
        clearAuth();
        goto('/');
    });
</script>

<div class="text-center py-8">
    <p>Logging out...</p>
</div> 