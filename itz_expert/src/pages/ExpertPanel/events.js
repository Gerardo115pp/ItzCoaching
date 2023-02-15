export const expert_panel_events = {
    EXPERT_DATA_CHANGED: 'EXPERT_DATA_CHANGED',
}

export const emitExpertDataChanged = () => {
    const event = new CustomEvent(expert_panel_events.EXPERT_DATA_CHANGED);
    document.dispatchEvent(event);
}