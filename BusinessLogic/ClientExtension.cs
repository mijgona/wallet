using DataAccess;

namespace BusinessLogic;

public static class UserExtension
{
    public static UserInfo ToUserInfo(this User client)
    {
        return new(
            Id: client.Id,
             UserName: client.UserName,
            Password: client.Password,
            FirstName: client.Name,
            LastName: client.LastName,
            MiddleName: client.MiddleName,
            Phone: client.PhoneNumber,
            PassportNumber: client.PassportNumber,
            Age: client.Age,
            Gender:  client.Gender.ToString());
    }
}
