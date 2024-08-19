import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RecruitManagementComponent } from './recruit-management.component';

describe('RecruitManagementComponent', () => {
  let component: RecruitManagementComponent;
  let fixture: ComponentFixture<RecruitManagementComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RecruitManagementComponent]
    });
    fixture = TestBed.createComponent(RecruitManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
