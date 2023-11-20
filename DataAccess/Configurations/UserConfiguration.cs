using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace DataAccess;

public sealed class UserConfiguration: IEntityTypeConfiguration<User>
{
    public void Configure(EntityTypeBuilder<User> modelBuilder)
    {
        modelBuilder
            .HasKey(t => t.Id)
            .HasName("pk_id");

        modelBuilder
            .Property(p => p.Id)
            .HasColumnType("SERIAL")
            .HasColumnName("id")
            .IsRequired();

        modelBuilder
            .HasOne(u => u.Balance)
            .WithOne(w => w.User)
            .HasForeignKey<WalletBalance>(w => w.UserId);
    }
}