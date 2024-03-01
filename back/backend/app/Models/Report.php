<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Report extends Model
{
    protected $fillable = ['user_id', 'appointment_id', 'procedure_id', 'price_id'];

    public function user()
    {
        return $this->belongsTo(User::class);
    }

    public function appointment()
    {
        return $this->belongsTo(Appointment::class);
    }

    public function procedure()
    {
        return $this->belongsTo(Procedure::class);
    }

    public function price()
    {
        return $this->belongsTo(Price::class);
    }
}
