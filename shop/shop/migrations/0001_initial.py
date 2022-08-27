# Generated by Django 4.1 on 2022-08-27 10:03

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Users',
            fields=[
                ('user_id', models.BigIntegerField(primary_key=True, serialize=False, verbose_name='ID пользователя')),
                ('user_name', models.TextField(max_length=40, verbose_name='@ пользователя')),
                ('first_name', models.TextField(max_length=40, null=True, verbose_name='Имя пользователя')),
                ('is_order_active', models.BooleanField(default=False, verbose_name='Сбор данных по заказу')),
            ],
            options={
                'verbose_name': 'Пользователь',
                'verbose_name_plural': 'Пользователи',
            },
        ),
        migrations.CreateModel(
            name='Orders',
            fields=[
                ('order_id', models.AutoField(primary_key=True, serialize=False, verbose_name='ID заказа')),
                ('order_value', models.CharField(max_length=200, verbose_name='Содержимое заказа')),
                ('status', models.SmallIntegerField(verbose_name='Статус заказа')),
                ('is_closed', models.BooleanField(default=False, verbose_name='Заказ выполнен')),
                ('date', models.DateTimeField(auto_now_add=True, verbose_name='Дата и время заказа')),
                ('is_chat_active', models.BooleanField(default=False, verbose_name='Активная переписка')),
                ('user_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='shop.users', verbose_name='ID пользователя')),
            ],
            options={
                'verbose_name': 'Заказ',
                'verbose_name_plural': 'Заказы',
            },
        ),
        migrations.CreateModel(
            name='Messages',
            fields=[
                ('message_id', models.BigAutoField(primary_key=True, serialize=False, verbose_name='ID сообщения')),
                ('message_text', models.CharField(max_length=200, verbose_name='Текст сообщения')),
                ('is_sender', models.BooleanField(verbose_name='Является ли пользователь отправителем?')),
                ('date', models.DateTimeField(auto_now_add=True, verbose_name='Дата и время отправки')),
                ('order_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='shop.orders', verbose_name='ID заказа')),
                ('user_id', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='shop.users', verbose_name='ID пользователя')),
            ],
            options={
                'verbose_name': 'Сообщение',
                'verbose_name_plural': 'Сообщения',
            },
        ),
    ]
