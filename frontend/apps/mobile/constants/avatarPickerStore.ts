export type AvatarPickerResult = { face: string | null; background: string };
type Callback = (result: AvatarPickerResult) => void;

let pendingCallback: Callback | null = null;

export function setPendingAvatarCallback(cb: Callback) {
  pendingCallback = cb;
}

export function resolvePendingAvatarCallback(result: AvatarPickerResult) {
  pendingCallback?.(result);
  pendingCallback = null;
}
