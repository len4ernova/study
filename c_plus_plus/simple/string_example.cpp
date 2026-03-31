import std;

int main(){
    std::string s = "The standart string class";

    const char c = s[1];

    const std::size_t n = s.size();

    s[n-1] = 'S';

    std::println("{}", s.empty);
    std::println("{} {}", c, n)

}