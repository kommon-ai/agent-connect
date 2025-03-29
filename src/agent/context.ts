/**
 * TypeScript implementation of a simple context, inspired by Go's context package.
 * Provides a way to signal cancellation to operations.
 */
export class Context {
  private cancelled: boolean = false;
  private timer: ReturnType<typeof setTimeout> | null = null;
  private cancelFns: (() => void)[] = [];
  
  /**
   * Creates a new Context with an optional timeout
   * @param timeoutMs - Optional timeout in milliseconds
   */
  constructor(timeoutMs?: number) {
    if (timeoutMs !== undefined && timeoutMs > 0) {
      this.timer = setTimeout(() => {
        this.cancel();
      }, timeoutMs);
    }
  }
  
  /**
   * Cancels the context and executes all registered cancel functions
   */
  cancel(): void {
    if (this.cancelled) return;
    
    this.cancelled = true;
    if (this.timer) {
      clearTimeout(this.timer);
      this.timer = null;
    }
    
    // Execute cancel functions in reverse order (LIFO)
    for (let i = this.cancelFns.length - 1; i >= 0; i--) {
      try {
        this.cancelFns[i]();
      } catch (error) {
        console.error('Error in cancel function:', error);
      }
    }
    this.cancelFns = [];
  }
  
  /**
   * Checks if the context has been cancelled
   * @returns true if the context is cancelled
   */
  isCancelled(): boolean {
    return this.cancelled;
  }
  
  /**
   * Register a function to be called when the context is cancelled
   * @param fn Function to call on cancellation
   * @returns Function to unregister the cancel function
   */
  onCancel(fn: () => void): () => void {
    if (this.cancelled) {
      fn(); // Execute immediately if already cancelled
      return () => {}; // Return no-op unregister function
    }
    
    this.cancelFns.push(fn);
    
    // Return a function to unregister this cancel function
    return () => {
      const index = this.cancelFns.indexOf(fn);
      if (index >= 0) {
        this.cancelFns.splice(index, 1);
      }
    };
  }
  
  /**
   * Creates a derived context that will be cancelled when the parent is cancelled
   * @param timeoutMs Optional timeout for the derived context
   * @returns A new derived context
   */
  withTimeout(timeoutMs: number): Context {
    const derivedContext = new Context(timeoutMs);
    
    // Link the derived context to this context
    const unregister = this.onCancel(() => {
      derivedContext.cancel();
      unregister(); // Clean up the registration
    });
    
    return derivedContext;
  }
  
  /**
   * Clean up resources when no longer needed
   */
  dispose(): void {
    this.cancel();
  }
}