using DataAccess;
using DataAccess.Repositories;

namespace BusinessLogic;

public sealed class UserService : IUserService
{
    private readonly IUserRepository _userRepository;

    public UserService(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public async Task<User> CreateUser(UserInfo userInfo)
    {
        User newUser = new()
        {
            Name = userInfo.FirstName,
            LastName = userInfo.LastName,
            MiddleName = userInfo.MiddleName,
            Age = userInfo.Age,
            PhoneNumber = userInfo.Phone,
            PassportNumber = userInfo.PassportNumber,
            Gender = userInfo.Gender.ToGenderEnum(),
            
        };

        return await _userRepository.CreateAsync(newUser);;
    }
    
    
    public async ValueTask<User> GetUserByUsername(string userName)
    {
        if (userName is not { Length: > 0 })
            throw new ArgumentNullException(nameof(userName));

        return await _userRepository.GetUserByUserName(userName, default);
    }

}
