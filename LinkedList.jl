import Base: peek

mutable struct LinkedList
    value::Any
    next::Union{LinkedList,Nothing}
end

function peek(list::LinkedList, n::Integer)
    out = Array{Int}(undef, n)
    for i in 1:n
        out[i] = list.value
        list = list.next
    end
    return out
end

@assert all(peek(LinkedList(1, LinkedList(2, LinkedList(3, LinkedList(4, nothing)))), 2) .== [1, 2])

function insertnext!(list::LinkedList, arr)
    head = list
    tail = list.next
    for x in arr
        head.next = LinkedList(x, nothing)
        head = head.next
    end
    head.next = tail
    return list
end

function cycle_length(list::LinkedList)
    len = 0
    start = list.value
    list = list.next
    while list.value != start
        len += 1
        list = list.next
    end
    return len
end