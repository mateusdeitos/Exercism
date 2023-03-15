type Node<T> = {
	value?: T;
	next?: Node<T>;
	prev?: Node<T>;
}

export class LinkedList<TElement> {
	private head: Node<TElement> = { value: undefined, next: undefined, prev: undefined };
	private tail: Node<TElement> = { value: undefined, next: undefined, prev: undefined };

	constructor() {
		this.head.next = this.tail;
		this.tail.prev = this.head;
	}

	public push(element: TElement) {
		const node: Node<TElement> = {
			value: element,
			next: this.tail,
			prev: this.tail?.prev,
		}

		if (node.prev) node.prev.next = node;
		if (node.next) node.next.prev = node;
	}

	public pop(): TElement {
		const node = this.tail.prev;
		this.tail.prev = node?.prev;
		if (node?.prev) node.prev.next = this.tail;
		return node?.value as TElement;
	}

	public shift(): TElement {
		const node = this.head.next;
		this.head.next = node?.next;
		if (node?.next) node.next.prev = this.head;
		return node?.value as TElement;
	}

	public unshift(element: TElement) {
		const node: Node<TElement> = {
			value: element,
			next: this.head?.next,
			prev: this.head,
		}

		if (node.prev) node.prev.next = node;
		if (node.next) node.next.prev = node;
	}

	public delete(element: TElement) {
		let node = this.head;
		do {
			if (node.value == element) {
				if (node.prev) node.prev.next = node.next;
				if (node.next) node.next.prev = node.prev;
				return;
			}

			if (!node.next) {
				break;
			}

			node = node.next;
		} while (node !== this.head);
	}

	public count(): number {
		let node = this.head;
		let count = 0;
		do {
			if (typeof node.value !== 'undefined') {
				count++;
			}

			if (!node.next) {
				break;
			}

			node = node.next;
		} while (node !== this.head);

		return count;
	}
}
