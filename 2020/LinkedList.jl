import Base: ==, peek, last

mutable struct LinkedList
    value::Any
    next::Union{LinkedList,Nothing}
end

function (==)(x::LinkedList, y::LinkedList)
    while !isnothing(x.next) && !isnothing(y.next)
        if x.value != y.value
            return false
        end
        x = x.next
        y = y.next
    end
    return x.value == y.value && isnothing(x.next) && isnothing(y.next)
end

@assert LinkedList(1, LinkedList(2, nothing)) == LinkedList(1, LinkedList(2, nothing))
@assert LinkedList(1, LinkedList(2, nothing)) != LinkedList(1, LinkedList(3, nothing))
@assert LinkedList(1, LinkedList(2, nothing)) != LinkedList(2, LinkedList(2, nothing))

function peek(list::LinkedList, n::Integer)
    out = Array{Int}(undef, n)
    for i in 1:n
        out[i] = list.value
        list = list.next
    end
    return out
end

@assert all(peek(LinkedList(1, LinkedList(2, LinkedList(3, LinkedList(4, nothing)))), 2) .== [1, 2])

function insertafter!(list::LinkedList, index, items::Vector)
    head = list
    while head.value != index
        head = head.next
    end

    tail = head.next
    for x in items
        head.next = LinkedList(x, tail)
        head = head.next
    end

    return list
end

@assert insertafter!(LinkedList(1, LinkedList(2, LinkedList(3, nothing))), 2, [7, 8]) ==
    LinkedList(1, LinkedList(2, LinkedList(7, LinkedList(8, LinkedList(3, nothing)))))

function tolinkedlist(arr::Vector)
    list = LinkedList(arr[end], nothing)
    for x in reverse(arr)[2:end]
        list = LinkedList(x, list)
    end
    return list
end

@assert tolinkedlist([1, 2, 3]) == LinkedList(1, LinkedList(2, LinkedList(3, nothing)))

function last(list::LinkedList)
    head = list
    while !isnothing(head.next)
        head = head.next
    end
    return head
end

@assert last(LinkedList(1, LinkedList(2, LinkedList(3, nothing)))) == LinkedList(3, nothing)
