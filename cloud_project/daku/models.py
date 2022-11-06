from django.db import models
from account.models import UserBase
# Create your models here.
class Container(models.Model):
    user = models.ForeignKey(UserBase, on_delete=models.CASCADE)
    container_id = models.CharField(max_length=200)

    class Meta:
        verbose_name = "Containers"
        verbose_name_plural = "Containers"

    def __str__(self):
        return self.container_id