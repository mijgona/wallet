namespace BusinessLogic;

public readonly record struct UserInfo(
    long Id,
    string UserName,
    string Password,
    string FirstName,
    string LastName,
    string? MiddleName,
    string Phone,
    string? PassportNumber,
    short Age,
    string Gender)
{
}

public record struct LoginRequest(
    
    string PhoneNumber,
    string Password)
{
}