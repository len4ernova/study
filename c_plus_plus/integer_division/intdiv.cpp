export module div;

import std;
std::size_t divide(std::size_t a, std::size_t b)
{
    if (b == 0){
        return std::numeric_limits<std::size_t>::max();
    }

    std::size_t res = 0;
    while (a >= b){
        std::size_t n = 0;
        std::size_t b_next = b << 1uz;

        while (a >= b_next && b < b_next){
            n++;
            b_next <<= 1uz;
        }
        a -= b << n;
        res += 1uz << n;
    }
    return res;
    
    
}