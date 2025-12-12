class AttendancesController < ApplicationController
  before_action :set_attendance, only: [:edit, :update, :delete, :destroy]

  def index
    @attendances = Attendance.all.order(created_at: :desc)
  end

  def new
    @attendance = Attendance.new
  end

  def create
    @attendance = Attendance.new(attendance_params)
    if @attendance.save
      redirect_to root_path, notice: '出社状況を登録しました'
    else
      render :new, status: :unprocessable_entity
    end
  end

  def edit
  end

  def update
    if @attendance.update(attendance_params)
      redirect_to root_path, notice: '出社状況を更新しました'
    else
      render :edit, status: :unprocessable_entity
    end
  end

  def delete
  end

  def destroy
    @attendance.destroy
    redirect_to root_path, notice: '出社状況を削除しました'
  end

  private

  def set_attendance
    @attendance = Attendance.find(params[:id])
  end

  def attendance_params
    params.require(:attendance).permit(:name, :location)
  end
end
