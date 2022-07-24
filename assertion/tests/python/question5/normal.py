# `mumble` is a function that accepts an argument as string and returns a
# string as its output.
def mumble(Input):
    result = []
    for index,letter in enumerate(Input):
        result.append( (letter*(index+1)).capitalize() )
    
    return "-".join(result)